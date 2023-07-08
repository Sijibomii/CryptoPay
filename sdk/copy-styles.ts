import fs from "node:fs"

fs.copyFile('src/main.css', 'dist/main.css', (error) => {
  if (error) {
    throw error;
  }
  // eslint-disable-next-line no-console
  console.log('main.css copied successfully.');
});