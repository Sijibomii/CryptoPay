import fs from "node:fs"

fs.copyFile('src/main.css', 'dist/esm/main.css', (error) => {
  if (error) {
    throw error;
  }
  // eslint-disable-next-line no-console
  console.log('main.css copied successfully.');
});


const directoryPath = 'dist/esm/images';

fs.mkdir(directoryPath, (err) => {
  if (err) {
    console.error('Failed to create directory:', err);
    return;
  }
  console.log('Directory created successfully.');
});

fs.copyFile('src/images/btc.png', 'dist/esm/images/btc.png', (error) => {
  if (error) {
    throw error;
  }
  // eslint-disable-next-line no-console
  console.log('btc.png copied successfully.');
});

fs.copyFile('src/images/error.png', 'dist/esm/images/error.png', (error) => {
  if (error) {
    throw error;
  }
  // eslint-disable-next-line no-console
  console.log('error.png copied successfully.');
});