'use strict';

const channels = require('./dormitories/client');

const client = channels.Client('http://localhost:8080');

(async () => {
    // Scenario 1: Get recommended dormitory for selected speciality.
    console.log('=== Scenario 1 ===');
    let dormitoryID;
    try {
        const { id, studentsCount } = await client.findDormitory('biology');
        dormitoryID = id;
        console.log('Recommended dormitory:', id);
        console.table(studentsCount);

    } catch (e) {
        console.log(` ${e.message}`);
    };

    // Scenario 2: Make record where the student was settled.
    console.log('=== Scenario 2 ===');
    try {
        const { id, specialty } = await client.commitSettle(dormitoryID, 'biology');
        console.log(`Student with specialty ${specialty} was settled in dormitory ${id}`)
    } catch (e) {
        console.log(` ${e.message}`);
    };
})();    