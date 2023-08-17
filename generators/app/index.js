const Generator = require('yeoman-generator');
const yosay = require('yosay');
const path = require('path');

const standalone = require('./generate-backend-standalone');
const microservice = require('./generate-microservice-monorepo');

const projectGenerators = [
    standalone, microservice
]

module.exports = class extends Generator{
    constructor(args, opts) {
        super(args, opts);
        this.description = 'Generates a gin backend project.';
        this.argument('destination', { type: String, required: false, description: `\n    The folder to create the project in, absolute or relative to the current working directory.\n    Use '.' for the current folder. If not provided, defaults to a folder with the display name.\n  ` })
        this.option('displayName', { type: String, alias: 'n', description: 'Display name of the project' });
        this.option('gitInit', { type: Boolean, description: `Initialize a git repo` });

        this.projectConfig = Object.create(null)
        this.projectGenerator = null;

        this.abort = false;
    }

    async initializing() {
        this.log(yosay('Welcome to the golang web backend generator!'));
        const destination = this.options['destination'];
        if (destination) {
            const folderPath = path.resolve(this.destinationPath(), destination);
            this.destinationRoot(folderPath);
        }
    }

    async prompting() {
        const choices = [];
        for (const g of projectGenerators) {
            const name = g.name;
            if (name) {
                choices.push({ name, value: g.id })
            }
        }
        this.projectConfig.type = (await this.prompt({
            type: 'list',
            name: 'type',
            message: 'What type of project do you want to create?',
            pageSize: choices.length,
            choices,
        })).type;

        this.projectGenerator = projectGenerators.find(g => g.id === this.projectConfig.type);
        try {
            // @ts-ignore
            await this.projectGenerator.prompting(this, this.projectConfig);
        } catch (e) {
            this.abort = true;
        }

    }

    async writing() {
        if (this.abort) {
            return;
        }
        if (!this.options['destination']) {
            this.destinationRoot(this.destinationPath(this.projectConfig.displayName))
        }
        this.env.cwd = this.destinationPath();

        this.log();
        this.log(`Writing in ${this.destinationPath()}...`);

        this.sourceRoot(path.join(__dirname, './templates/' + this.projectConfig.type));
        
        // @ts-ignore
        return this.projectGenerator.writing(this, this.projectConfig);
    }

    // todo: install swag, exec swag init, go get
    // async install() {
        
    // }

    async end() {
        if (this.abort) {
            return;
        }
        // Git init
        if (this.projectConfig.gitInit) {
            this.spawnCommand('git', ['init', '--quiet']);
        }

    }
}
