import React, { Component } from 'react'
import "./index.css"
import * as QrCode from 'qrcode.react'
import {getAuthUrl} from "../../request/auth"
import httpCode from "../../conventions/httpCode"


class Login extends Component {

    constructor(props) {
        super(props)
        this.state = {url:""}
    }

    componentDidMount(){
        getAuthUrl(new Date().getTime() + Math.random()).then((res)=>{
            const {status,data:{data:{url},message}} = res
            console.log(url)
            if(status === httpCode.OK) {
                this.setState({
                    url,
                })
            } else {
                alert(message)
            }
        })
    }

    render () {
        return (
            <div className="login_container">
                <div className="qrcode_box">
                    <h3>微信登录</h3>
                    <QrCode value={this.state.url} size={280} />
                    <div className="description_box">
                        <p className="d1">请用微信扫描二维码</p>
                        <p className="d2">“登录yry在线聊天室”</p>
                    </div>
                </div>
            </div>
        )
    }
}

export default Login