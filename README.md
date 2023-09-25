[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <img src="images/logo.png" alt="Logo" width="315" height="100">

  <h3 align="center">Easy Wallet - API</h3>

  <p align="center">
    This is the back end of the Easy Wallet project, a web application that helps you to manage your finances.
  <br />
    <strong>Status: <u>Development</u>

  </strong>
    <br />
    <a href="#"><strong>Explore the docs »</strong></a>
    <br />
    <a href="https://easywallet2023.netlify.app/">View Demo</a>
    ·
    <a href="https://github.com/TiagoRibeiro25/Easy-Wallet-BE/issues">Report Bug</a>
    ·
    <a href="https://github.com/TiagoRibeiro25/Easy-Wallet-BE/issues">Request Feature</a>
  </p>
   <img src="images/devices.png" alt="Logo" >
 <br />
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
       <li><a href="#prepare-for-production">Prepare for production</a></li>
      </ul>
    </li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About The Project

Easy Wallet is a web application that allows users to manage their finances. It was made with the purpose of helping people to organize their money and to have a better understanding of their expenses and incomes.

This project was made as a side project with the purpose of getting familiar with the technologies used in the development of the application.

### Built With

* [Golang](https://golang.org/)
* [Echo](https://echo.labstack.com/)
* [GORM](https://gorm.io/)
* [PostgreSQL](https://www.postgresql.org/)
* [Docker](https://www.docker.com/)

<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

In order to run this project, you need to have installed

* Golang*
* Nodejs and npm (in case you want to use [nodemon](https://www.npmjs.com/package/nodemon))
* Docker and docker-compose (in case you want to use the provided docker-compose file to run the PostgreSQL database)

You also need to have access to the a:

* PostgreSQL database (or use the provided docker-compose file)
* Redis database (or use the provided docker-compose file)
* Cloudinary account
* Mailjet account

### Installation

1. Clone the repo

   ```sh
   git clone https://github.com/TiagoRibeiro25/Easy-Wallet-BE.git
   ```

2. Install NPM packages (in case you want to use nodemon)

3. Install Go packages

    ```sh
   go mod download
   ```

4. Create a .env file, add the variables in the .env.example file and fill them with your own values.

5. Have a PostgreSQL database and Redis database running (or use the provided docker-compose file)

6. Run the project

* Using nodemon

   ```sh
   npm run dev
   ```

* Using go run

   ```sh
   go run ./src
   ```

<!-- Production -->
## Prepare for production

1. Run the go mod tidy command to install the dependencies

   ```sh
   go mod download
   ```

2. Build the project

* Using npm

   ```sh
   npm run build
   ```

* Using go build

  ```sh
  go build -o ./dist/easywalletapi ./src
  ```

3. Set the GO_ENV variable to production and the rest of the variables in the .env file

4. Run the project

   ```sh
   ./dist/easywalletapi
   ```

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<!-- CONTACT -->
## Contact

Contact through my [personal website](https://tiagoribeiro.tech/contact).

<!-- MARKDOWN LINKS & IMAGES -->
[contributors-shield]: https://img.shields.io/github/contributors/TiagoRibeiro25/Easy-Wallet-BE.svg?style=for-the-badge
[contributors-url]: https://github.com/TiagoRibeiro25/Easy-Wallet-BE/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/TiagoRibeiro25/Easy-Wallet-BE.svg?style=for-the-badge
[forks-url]: https://github.com/TiagoRibeiro25/Easy-Wallet-BE/network/members
[stars-shield]: https://img.shields.io/github/stars/TiagoRibeiro25/Easy-Wallet-BE.svg?style=for-the-badge
[stars-url]: https://github.com/TiagoRibeiro25/Easy-Wallet-BE/stargazers
[issues-shield]: https://img.shields.io/github/issues/TiagoRibeiro25/Easy-Wallet-BE.svg?style=for-the-badge
[issues-url]: https://github.com/TiagoRibeiro25/Easy-Wallet-BE/issues
[license-shield]: https://img.shields.io/github/license/TiagoRibeiro25/Easy-Wallet-BE.svg?style=for-the-badge
[license-url]: https://github.com/TiagoRibeiro25/Easy-Wallet-BE/blob/master/LICENSE.txt
