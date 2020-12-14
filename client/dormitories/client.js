
'use strict'

const http = require('../common/http');

const Client = (baseUrl) => {

    const client = http.Client(baseUrl);

    return {
        findDormitory: (specialty) => client.get('/dormitories?specialty='+ specialty),
        commitSettle: (dormitoryID, specialty) => client.post('/dormitories/', {dormitoryID, specialty})
    }

};

module.exports = { Client };


