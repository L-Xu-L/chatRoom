import React, { Component } from 'react'
import "./index.css"
class Chat extends Component {
    render () {
        return (
            <div className="chat_container">
                <div className="content">
                    <div className="panel">
                        <div className="personal">
                            <img src="#" alt="我的头像"/>
                            <p className="nickName">帅羊</p>
                        </div>
                        <div className="user-order">
                            <div className="user-item">
                                <img src="#" alt="用户头像"/>
                                <p className="nickName">马云</p>
                                <p className="last-time">17:35</p>
                            </div>
                        </div>
                    </div>
                    <div className="window">
                        <div className="top"></div>
                        <div className="content"></div>
                        <div className="message">
                            <div className="toolBar">

                            </div>
                            <div className="message-body">

                            </div>
                            <div className="bottom">
                                <button>发送</button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default Chat