const path = require('path');

exports.askForExtensionDisplayName = (generator, config) => {
    let extensionDisplayName = generator.options['extensionDisplayName'];
    if (extensionDisplayName) {
        config.displayName = extensionDisplayName;
        return Promise.resolve();
    }
    const nameFromFolder = generator.options['destination'] ? path.basename(generator.destinationPath()) : 'helloWorld';

    if (generator.options['quick'] && nameFromFolder) {
        config.displayName = nameFromFolder;
        return Promise.resolve();
    }

    return generator.prompt({
        type: 'input',
        name: 'displayName',
        message: 'What\'s the name of your project?',
        default: nameFromFolder
    }).then(displayNameAnswer => {
        config.displayName = displayNameAnswer.displayName;
    });
}

exports.askForGit = (generator, config) => {
    let gitInit = generator.options['gitInit'];
    if (typeof gitInit === 'boolean') {
        config.gitInit = Boolean(gitInit);
        return Promise.resolve();
    }
    if (generator.options['quick']) {
        config.gitInit = true;
        return Promise.resolve();
    }

    return generator.prompt({
        type: 'confirm',
        name: 'gitInit',
        message: 'Initialize a git repository?',
        default: true
    }).then(gitAnswer => {
        config.gitInit = gitAnswer.gitInit;
    });
}