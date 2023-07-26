const fs = require('fs');
const path = require('path');

module.exports = {
  reactStrictMode: false,
  server: {
    https: {
      key: fs.readFileSync(path.resolve('../server-app/certificate/localhost.key')),
      cert: fs.readFileSync(path.resolve('../server-app/certificate/localhost.crt')),
    },
  },
};
