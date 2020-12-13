'use strict';

const channels = require('./channels/client');

const client = channels.Client('http://localhost:8080');

// Scenario 1: Get recommended dormitory for selected speciality.
client.findDormitory()
    .then(() => {
        console.log('=== Scenario 1 ===');
        console.log('Recommended dormitory:');

    })
    .catch((e) => {
        console.log(` ${e.message}`);
    });

// Scenario 2: Make record where the student was settled.
client.createChannel()
    .then(() => {
        console.log('=== Scenario 2 ===');
        console.log('');
    })
    .catch((e) => {
        console.log(` ${e.message}`);
    });