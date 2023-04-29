import axios from "axios";

export default axios.create({
    baseURL: "http://go_app:4000"
})