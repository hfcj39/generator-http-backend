# generator-http-backend
## usage
Since you’re developing the generator locally, it’s not yet available as a global npm module. A global module may be created and symlinked to a local one, using npm.
```bash
git clone https://github.com/hfcj39/generator-http-backend.git
cd generator-http-backend
npm link
```
then you'll be able to call 
```
yo http-backend
```
