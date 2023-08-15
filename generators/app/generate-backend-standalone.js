const prompts = require("./prompts");

module.exports = {
    id: 'backend-standalone',
    name: 'Standalone web app backend',
    prompting: async (generator, config) => {
        generator.log(config)
        await prompts.askForExtensionDisplayName(generator, config)
        generator.log(config)
    },
    writing: async (generator, config) => {
        console.log(config)
        generator.fs.copyTpl(generator.templatePath('README.md'), generator.destinationPath('README.md'), config);
        generator.fs.copyTpl(generator.templatePath(''), generator.destinationPath(''), config);
    }

}