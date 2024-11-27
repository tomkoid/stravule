import { PUBLIC_BACKEND_URL } from "$env/static/public"
import { getUserInfo } from "./user_info"

interface LoginResponse {
  canteen: number,
  sid: number
  user_hash: string
}
// returns if logged in
export async function login(username: string, password: string, canteen: string) {
  let params = new URLSearchParams()

  params.append("username", username)
  params.append("password", password)
  params.append("canteen", canteen)
  let req = await fetch(`${PUBLIC_BACKEND_URL}/api/v1/login?${params.toString()}`, {
    method: "POST",
  })

  if (req.status != 200) {
    const errorText = await req.text()
    throw new Error(errorText)
  }

  let data: LoginResponse = await req.json()
  console.log(data)

  localStorage.setItem("sid", data.sid.toString())
  localStorage.setItem("canteen", data.canteen.toString())
  localStorage.setItem("user_hash", data.user_hash)

  localStorage.setItem("username", username)
  localStorage.setItem("password", password)

  await getUserInfo(data.sid.toString(), data.canteen.toString())
}
