
'use strict'

const http = require('../common/http');

const Client = (baseUrl) => {

    const client = http.Client(baseUrl);

    return {
        findDormitory: () => client.get('/dormitories'),
        commitSettle: () => client.post('/dormitories', {dormitoryID, speciality})
    }

};

module.exports = { Client };


