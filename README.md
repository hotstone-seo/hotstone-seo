# HotStone Center

Dashboard for HotStone-SEO

## Configuration

Update configuration at `.env` file. 
```
PORT=3000
CHOKIDAR_USEPOLLING=true
```

Please use `.env.template` if no `.env` available

---

## Development Guide

### Prerequisite

Install [NodeJS](https://nodejs.org/en/)
```bash
brew install node
```

### Build

```bash
npm install

# start the server
npm start
```

### Editor

It's recommended to using [Visual Studio Code](https://code.visualstudio.com/) as editor for this project

Extensions:
- (Syntax Highlight)[https://marketplace.visualstudio.com/items?itemName=mgmcdermott.vscode-language-babel]
- (Code Formatter)[https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode]

Settings (`cmd`+`shift`+`p`, type "Preferences:Open Settings (JSON)"):
```json
{
    // Your current settings

    "[javascriptreact]": {
        "editor.formatOnSave": true,
        "editor.defaultFormatter": "esbenp.prettier-vscode"
    },
    "[javascript]": {
        "editor.formatOnSave": true,
        "editor.defaultFormatter": "esbenp.prettier-vscode"
    }
}
```


