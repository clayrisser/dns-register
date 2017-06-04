# dns-register

Automatically register and unregister your server with CloudFlare DNS

Please &#9733; this repo if you found it useful &#9733; &#9733; &#9733;


## Features
<!------------------------------------------------------->

* Automatically register server's public IP as an A record
* Automatically unregister server's public IP as an A record


## Installation
<!------------------------------------------------------->

```sh
git clone https://github.com/jamrizzi/dns-register
cd dns-register
export GOPATH=$(pwd)
make dns-register
sudo mv ./dns-register /bin/dns-register
```


## Dependencies
<!------------------------------------------------------->

* [GoLang](https://golang.org/)
* [Make](https://www.gnu.org/software/make/)
* [Docker](https://www.docker.com/)


## Usage
<!------------------------------------------------------->

### Command Line Tool

Set Environment Variables

```sh
export CLOUDFLARE_API_KEY=your-cloudflare-api-key
export CLOUDFLARE_EMAIL=your-cloudflare-email
export CLOUDFLARE_WEBSITE=your-cloudflare-website
export SUBDOMAIN=servers
```

* Register
  
  ```sh
  dns-register register
  ```
  
* Unregister
  
  ```sh
  dns-register unregister
  ```
  
### Docker

  * Register
  
  ```sh
  docker run --rm \
    -e CLOUDFLARE_API_KEY=your-cloudflare-api-key \
    -e CLOUDFLARE_EMAIL=your-cloudflare-email \
    -e CLOUDFLARE_WEBSITE=your-cloudflare-website \
    -e SUBDOMAIN=servers -e \
    jamrizzi/dns-register:latest register
  ```

  * Unregister
  
    ```sh
    docker run --rm \
      -e CLOUDFLARE_API_KEY=your-cloudflare-api-key \
      -e CLOUDFLARE_EMAIL=your-cloudflare-email \
      -e CLOUDFLARE_WEBSITE=your-cloudflare-website \
      -e SUBDOMAIN=servers -e \
      jamrizzi/dns-register:latest unregister
    ```


## Support
<!------------------------------------------------------->

Submit an [issue](https://github.com/jamrizzi/dns-register/issues/new)


## Buy Me Coffee
<!------------------------------------------------------->

A ridiculous amount of coffee was consumed in the process of building this project.

[Add some fuel](https://pay.jamrizzi.com) if you'd like to keep me going!


## Contributing
<!------------------------------------------------------->

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -m 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D


## License
<!------------------------------------------------------->

[MIT License](https://github.com/jamrizzi/dns-register/blob/master/LICENSE)

[Jam Risser](https://jamrizzi.com) &copy; 2017


## Credits
<!------------------------------------------------------->

* [Jam Risser](https://jamrizzi.com) - Author


## Changelog
<!------------------------------------------------------->

0.0.1 (2017-06-03)
* Initial release
