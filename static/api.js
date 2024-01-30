async function getPackSettings() {
    const url = 'http://localhost:8080/submitPackSettings';

    try {
        const response = await fetch(url, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
        });

        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
        }

        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Error:', error);
        return null; // or handle the error as needed
    }
}



async function submitPackSettings(settings) {
    const url = 'http://localhost:8080/submitPackSettings';

    try {
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(settings),
        });

        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
        }

        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Error:', error);
        return null; // or handle the error as needed
    }
}

// Example usage
const packSettings = [23, 31, 53];

async function performSubmitPackSettings() {
    try {
        const response = await submitPackSettings(packSettings);
        console.log('Response:', response);
    } catch (error) {
        console.error('Error:', error);
    }
}

// // Call the async function
// performSubmitPackSettings();


async function calculatePacks(totalNumberOfPacks) {
    const url = 'http://localhost:8080/calculatePacks';

    try {
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                totalNumberOfPacks: totalNumberOfPacks,
            }),
        });

        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
        }

        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Error:', error);
        return null; // or handle the error as needed
    }
}

// Example usage
const totalNumberOfPacks = 263;

async function performCalculatePacks() {
    try {
        const response = await calculatePacks(totalNumberOfPacks);
        console.log('Response:', response);
    } catch (error) {
        console.error('Error:', error);
    }
}

// Call the async function
// performCalculatePacks();


document.addEventListener("DOMContentLoaded", function (event) {
    // Your code to run since DOM is loaded and ready

    getPackSettings().then(response => {
        console.log('Response:', response);
    });

    
});