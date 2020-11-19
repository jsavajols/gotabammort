var publiques = require('../publiques/publiques');
var chaines = require('../global/chaines');
var logs = require('../global/logs');
var moment = require('moment');

module.exports = {
    moteur: function moteur(capital, duree, txInteret, detail, dateDebut, dateEffet, mode, assurance, mandat) {
        return new Promise(function (resolve, reject) {
            var dateEcheance = moment(dateDebut);
            var txInteretOri = txInteret;
            txInteret = (txInteret / 100) / 12;
            var ta = [];
            var periode;
            var interets;
            var mensu;
            var crdDebutMen = capital;
            var crdDebutAn = capital;
            var crdFin;
            var annee = 1;
            var cumulInterets = 0;
            var crdSit = 0;
            var dureeRest = 0;
            if (txInteret === 0) {
                mensu = (capital / duree);
            } else {
                mensu = (capital * txInteret) / (1 - Math.pow((1 + (txInteret)), -duree));
            }
            for (var m = 1; m <= duree; m++) {
                interets = crdDebutMen * txInteret;
                crdFin = crdDebutMen - mensu + interets;
                if (detail === "d") {
                    periode = {
                        mois: m,
                        dateEcheance: dateEcheance,
                        crdDebut: crdDebutMen,
                        crdFin: crdFin,
                        mensualite: mensu,
                        interets: interets,
                    };
                    // Regarde à quel endroit du TA on se situe 
                    if (moment(dateEcheance).year() === moment(dateEffet).year() &&
                        moment(dateEcheance).month() === moment(dateEffet).month()) {
                        crdSit = Math.trunc(crdDebutMen);
                        if (m === 1) {
                            dureeRest = duree;
                        } else {
                            dureeRest = (duree - m);
                        }
                    }
                    dateEcheance = moment(dateEcheance).add(1, "month").format("YYYY-MM-DD");
                    ta.push(periode);
                } else {
                    if (m / 12 === Math.trunc(m / 12)) {
                        periode = {
                            annee: annee,
                            crdDebut: crdDebutAn,
                            crdFin: crdFin,
                            mensualite: mensu,
                            interets: cumulInterets,
                        };
                        ta.push(periode);
                        cumulInterets = 0;
                        crdDebutAn = crdFin;
                        annee++;
                    }
                }
                crdDebutMen = crdFin;
                cumulInterets = cumulInterets + interets;
            }

            if (mode === 'crd') {
                resolve({
                    capitalInitial: capital,
                    dureeInitiale: duree,
                    dateDebutPret: dateDebut,
                    dateEffet: dateEffet,
                    crd: crdSit,
                    dureeRest: dureeRest,
                });
            } else {

                // Calcul du TEG et TAEA
                if ((capital + cumulInterets + assurance + mandat) < capital) {
                    return -1;
                }
                var tauxmin = 0;
                var tauxmax = 25;
                var taux = tauxmin;
                var taux_pour_calcul;
                var cout_total = 0;
                var coutTotal = (capital + cumulInterets + assurance + mandat);
                while (Math.abs(((tauxmax - tauxmin) * 1000000 > 0.00001))) {
                    taux = ((tauxmax + tauxmin) / 2);
                    taux_pour_calcul = (Math.pow((1 + (taux / 100)), (1 / 12))) - 1;
                    cout_total = mt_echeance(capital, taux_pour_calcul, duree) * duree;
                    if (cout_total < coutTotal) {
                        tauxmin = taux;
                    } else {
                        tauxmax = taux;
                    }
                }
                var taea = taux - txInteretOri;
                taux = Number.parseFloat(taux).toFixed(2);
                function mt_echeance(capital, taux_pour_calcul, duree) {
                    var val = Math.pow(taux_pour_calcul + 1, duree);
                    var eche = ((capital * taux_pour_calcul) * val) / (val - 1);
                    return eche;
                }

                resolve({
                    capitalInitial: capital,
                    dureeInitiale: duree,
                    dateDebutPret: dateDebut,
                    dateEffet: dateEffet,
                    crd: crdSit,
                    mensu: Number.parseFloat(mensu).toFixed(2),
                    dureeRest: dureeRest,
                    cumulInterets: Number.parseFloat(cumulInterets).toFixed(2),
                    assurance: assurance,
                    coutGlobal: Number.parseFloat(cumulInterets + assurance + mandat).toFixed(2),
                    taeg: Number.parseFloat(taux).toFixed(4),
                    taea: Number.parseFloat(taea).toFixed(4),
                    txmoyen: Number.parseFloat((((assurance / capital) / duree) * 12) * 100).toFixed(4),
                    ta: ta
                });

            }
        });
    }
}

/*
async function start() {
    var moteur = require('../moteurCalcul/moteur');
    var myTa = await moteur.moteur(130000, 300, 1.34, "d", "2000-01-01", "2020-08-01", "", 3555.68, 690);
    logs.log("TAEG        : " + myTa.taeg);
    logs.log("TAEA        : " + myTa.taea);
    logs.log("TX MOYEN    : " + myTa.txmoyen);
    logs.log("Mensualité  : " + myTa.mensu);
    logs.log("Intérêts    : " + myTa.cumulInterets);
    logs.log("Assurance   : " + myTa.assurance);
    logs.log("Coût Global : " + myTa.coutGlobal);
}

start();
*/
