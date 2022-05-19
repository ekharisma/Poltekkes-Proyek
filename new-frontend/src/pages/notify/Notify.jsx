import React from 'react';
import Sidebar from '../../components/sidebar/Sidebar';
import Navbar from "../../components/navbar/Navbar"
import "./notify.css"

const Notify = () => {
    return (
        <>
            <Navbar />
            <div className="notify">
                <Sidebar />
                <div className="notifyContainer">
                    <div className="qrcodeContainer">
                        <img src="/images/qrcode.svg" className='qrcode' alt="" />
                    </div>
                    <p className='notifyLabel'>Scan qrcode diatas untuk terhubung dengan bot VaccineBox</p>
                    {/* <Selector /> */}

                </div>
            </div>
        </>

    )
}

export default Notify