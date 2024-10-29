interface User {
  cislo: string,
  dochazka: string,
  editaceJidelnicku: number,
  email: string,
  id: string,
  jmeno: string,
  konto: string,
  limit: string,
  mena: string,
  nazevJidelny: string,
  pasivni: boolean,
  pocetJidel: boolean,
  podrobnosti: boolean,
  prihlaska: boolean,
  slabeHeslo: boolean,
  text: string,
  verze: string,
  vicenasobny: boolean,
  vs: string,
  vydej: number,
}

interface LoginResponse {
  betatest: boolean,
  cislo: string,
  ignoreCert: boolean,
  jmeno: string,
  s5url: string,
  sid: string,
  uzivatel: User,
  verze: string,
  zustatPrihlasen: true,
}
// returns if logged in
export async function login(username: string, password: string, canteen: string) {
  let req = await fetch("https://app.strava.cz/api/login", {
    method: "POST",
    referrer: "https://app.strava.cz/en/prihlasit-se?jidelna",
    referrerPolicy: "strict-origin-when-cross-origin",
    headers: {
      "accept": "*/*",
      "accept-language": "en-US,en;q=0.9",
      "content-type": "text/plain;charset=UTF-8",
      "priority": "u=1, i",
      "sec-ch-ua": "\"Not?A_Brand\";v=\"99\", \"Chromium\";v=\"130\"",
      "sec-ch-ua-mobile": "?0",
      "sec-ch-ua-platform": "\"Linux\"",
      "sec-fetch-dest": "empty",
      "sec-fetch-mode": "cors",
      "sec-fetch-site": "same-origin",
      "sec-gpc": "1"
    },
    mode: "no-cors",
    body: JSON.stringify({
      cislo: canteen,
      jmeno: username,
      heslo: password,
      zustatPrihlasen: true,
      environment: "W",
      lang: "EN"
    })
  })

  if (req.status != 200) {
    throw new Error(await req.text())
  }

  let data: LoginResponse = await req.json()
  console.log(data)
}
