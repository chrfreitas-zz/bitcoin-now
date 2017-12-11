const axios = require('axios')
const console = require('console')

const url = 'https://blockchain.info/pt/ticker',
      currency = process.argv[2];

axios.get(url).then(({ data }) => {
    if(!currency) {
        throw new Error('There is not parameters')
    }

    if(!data[currency]) {
        throw new Error('This currency doens\'t exist')
    }

    console.log(data[currency].last);
})
