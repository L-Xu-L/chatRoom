import {auth} from "../conventions/url"
import axios from "axios"

//获取微信授权url
const getAuthUrl = (state) => axios.get(`${auth}/url?state=${state}`)
export {
    getAuthUrl
}