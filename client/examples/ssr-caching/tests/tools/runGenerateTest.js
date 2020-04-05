const path = require('path'),
  fs = require('fs');
const { generateFile } = require('./generateTest');

let counter = 0;

function fromDir(startPath, fileExt) {
  if (!fs.existsSync(startPath)) {
    console.log('no dir ', startPath);

    return Promise.reject();
  }

  const files = fs.readdirSync(startPath);

  return Promise.all(
    files.map(file => {
      const filename = file;
      const filePath = path.join(startPath, file);
      const stat = fs.lstatSync(filePath);

      if (stat.isDirectory()) {
        return fromDir(filePath, fileExt); //recurse
      } else if (filename.indexOf('.spec') === -1 && filename.indexOf(fileExt) >= 0) {
        let paths = filePath.split('/');
        const name = filename.indexOf('index') === -1 ? filename : paths[paths.length - 2];

        return generateFile(`./${filePath}`, name, filename, startPath).then(
          numberOfCreatedFiles => (counter = counter + numberOfCreatedFiles)
        );
      } else {
        return Promise.resolve();
      }
    })
  );
}

const projectPath = process.argv[2];

if (!projectPath) {
  throw Error('Not path given');
}
const ext = process.argv[3] || '.js';

fromDir(projectPath, ext)
  .then(() => {
    console.log(`Finished, created ${counter} jest snapshot files`);
  })
  .catch(err => console.log('neivel', err));
