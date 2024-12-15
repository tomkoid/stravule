<a id="readme-top"></a>

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://codeberg.org/tomkoid/stravule">
    <img src="frontend/static/favicon.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Stravule</h3>

  <p align="center">
    Automatický vybírač obědů pro <a href="https://strava.cz">strava.cz</a>
    <br />
    <a href="https://github.com/othneildrew/Best-README-Template"><strong>Oficiální stránka »</strong></a>
    <br />
    <br />
    <a href="https://codeberg.org/tomkoid/stravule/issues/new">Nahlásit Bug</a>
    ·
    <a href="https://codeberg.org/tomkoid/stravule/issues/new">Požadavek o novou funkci</a>
  </p>
</div>

<!-- ABOUT THE PROJECT -->
## O projektu 

<!-- [![Stravule Showcase][stravule-showcase]](https://example.com) -->

*Upozornění: Stravule není jakýmkoli způsobem ovlivněna či spojena se společností VIS Plzeň*.

Stravuli jsem vytvořil jako řešení situace ve škole, kde jsou pro mě často méně oblíbená jídla automaticky nastaveny. Pomocí Stravule si mohu nastavit pozitivní a negativní filtry. Pokud-li název objednávky obsahuje klíčové slovo (filtr), nastaví se podle něj.

Například když mám pozitivní filtr *buchtičky* a negativní filtr *brambory* a mám na výběr ze dvou možností:
1. ***Špenát, vejce, brambory*** 
2. ***Buchtičky se šodó***

Objedná se ***Buchtičky se šodó***, protože obsahuje klíčové slovo *buchtičky*.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Vyrobeno pomocí

* ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
* ![TypeScript](https://img.shields.io/badge/typescript-%23007ACC.svg?style=for-the-badge&logo=typescript&logoColor=white)
* ![Svelte](https://img.shields.io/badge/svelte-%23f1413d.svg?style=for-the-badge&logo=svelte&logoColor=white)
* ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white) 
* ![Git](https://img.shields.io/badge/git-%23F05033.svg?style=for-the-badge&logo=git&logoColor=white)
* [![forthebadge](https://forthebadge.com/images/badges/license-mit.svg)](https://forthebadge.com)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Jak spustit?

### docker-compose

Tato metoda se hodí k hostování Stravule či k rychlému demu bez větší námahy.

Předtím, než začnete, ujistěte se, že máte nainstalovaný Docker společně s docker-compose.

Poté stačí pouze jít do naklonované složky a spustit:

```bash
docker compose up --build
```

Nyní otevřete webový prohlížeč s URL `http://localhost:8088`. Možnosti a networking je konfigurovatelný v `compose.yml`.

### Manuálně (určené pro vývojáře)

*Napsané pomocí AI*

Ujistěte se, že máte nainstalované potřebné nástroje:

- Go (pro backend)
- Node.js a npm (pro frontend)

Otevřete terminál spustťe tyto příkazy.

1. Nejprve se postarejte o spuštění backendu`:

   ```bash
   cd backend/
   go build . -o stravule
   ./stravule
   ```

2. Poté spustíme v druhém terminálu frontend:
   ```bash
   cd frontend/
   npm install
   npm run dev
   ```

<!-- USAGE EXAMPLES -->
<!-- ## Usage

Use this space to show useful examples of how a project can be used. Additional screenshots, code examples and demos work well in this space. You may also link to more resources.

_For more examples, please refer to the [Documentation](https://example.com)_

<p align="right">(<a href="#readme-top">back to top</a>)</p> -->



<!-- ROADMAP -->
<!-- ## Roadmap

- [x] Add Changelog
- [x] Add back to top links
- [ ] Add Additional Templates w/ Examples
- [ ] Add "components" document to easily copy & paste sections of the readme
- [ ] Multi-language Support
    - [ ] Chinese
    - [ ] Spanish -->

<!-- See the [open issues](https://github.com/othneildrew/Best-README-Template/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p> -->



<!-- CONTRIBUTING -->
## Vývoj 

Jakýkoliv váš příspěvek je **velmi ceněn**.

Máte-li návrh, jak tento projekt vylepšit, udělejte fork tohoto repozitáře a pošlete pull request.

1. Forkněte projekt
2. Vytvořte si git branch pro svou funkci (`git checkout -b feature/coolfeature`)
3. Proveďte commit svých změn (`git commit -m 'feat: add cool feature'`)
4. Pushněte změny do branche (`git push origin feature/coolfeature`)
5. Otevřete Pull Request

<!-- LICENSE -->
## Licence

Distributováno pod MIT licencí. Více informací v souboru `LICENSE`.

<p align="right">(<a href="#readme-top">back to top</a>)</p>
