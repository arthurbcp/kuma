{
  "name": "{{ .global.projectName }}",
  "version": "1.0.0",
  "description": "",
  "main": "dist/index.js",
  "types": "dist/index.d.ts",
  "scripts": {
    "build": "tsc",
    "watch": "tsc -w",
    "test": "jest",
    "lint": "eslint 'src/**/*.{ts,tsx}'",
    "format": "prettier --write 'src/**/*.{ts,tsx}'",
    "prepare": "npm run build"
  },
  "repository": {
    "type": "git",
    "url": ""
  },
  "keywords": [
    "typescript",
    "library"
  ],
  "author": "",
  "license": "",
  "bugs": {
    "url": "{{ .global.projectRepository }}/issues"
  },
  "homepage": "{{ .global.projectName }}#readme",
  "devDependencies": {
    "@types/jest": "^29.0.0",
    "@typescript-eslint/eslint-plugin": "^5.0.0",
    "@typescript-eslint/parser": "^5.0.0",
    "eslint": "^8.0.0",
    "jest": "^29.0.0",
    "prettier": "^2.0.0",
    "ts-jest": "^29.0.0",
    "typescript": "^5.0.0"
  },
  "peerDependencies": {
    "react": "^18.0.0"
  },
  "dependencies": {
    "axios": "^1.7.7"
  }
}
