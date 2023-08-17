const prompts = require("./prompts");

module.exports = {
    id: 'backend-standalone',
    name: 'Standalone web app backend',
    prompting: async (generator, config) => {
        generator.log(config)
        await prompts.askForExtensionDisplayName(generator, config)
        await prompts.askForGit(generator, config)
        generator.log(config)
    },
    writing: async (generator, config) => {
        console.log(config)
        generator.fs.copyTpl(generator.templatePath('README.md'), generator.destinationPath('README.md'), config);
        generator.fs.copyTpl(generator.templatePath(''), generator.destinationPath(''), config);
        generator.fs.copy(generator.templatePath('.air.toml'), generator.destinationPath('.air.toml'));
        generator.fs.copy(generator.templatePath('.mailmap'), generator.destinationPath('.mailmap'));
        generator.fs.copy(generator.templatePath('.dockerignore'), generator.destinationPath('.dockerignore'));
        if (config.gitInit) {
            generator.fs.copy(generator.templatePath('gitignore'), generator.destinationPath('.gitignore'));
        }
    }

}