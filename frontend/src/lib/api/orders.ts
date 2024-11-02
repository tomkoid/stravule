export interface Order {
  datum: string,
  druh: string,
  chod: string,
  pocet: number,
  cena: string,
  omezeni: string,
  polevka: string,
  id: number,
  nazev: string,
  druhPopis: string,
  druhChod: string,
  databaze: string,
  delsiPopis: string,
  alergenyText: string,
  zkratkaProduktu: string,
  multipleNazev: string,
  version: number,
  casKonec: string,
  casOdhlaseni: string,
}
// returns if logged in
export async function loadOrders(sid: string, canteen: string, pickOrders: boolean): Promise<Order[][]> {
  let params = new URLSearchParams()

  params.append("sid", sid)
  params.append("canteen", canteen)
  params.append("pick", Boolean(pickOrders).toString())
  let req = await fetch(`http://127.0.0.1:1323/api/v1/orders?${params.toString()}`, {
    method: "POST",
  })

  if (req.status != 200) {
    throw new Error(await req.text())
  }

  let data: Order[][] = await req.json()

  return data
}
