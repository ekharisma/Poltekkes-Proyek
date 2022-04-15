import React from 'react';
import Navbar from "../../components/navbar/Navbar"
import Selector from "../../components/selector/Selector"
import Sidebar from "../../components/sidebar/Sidebar"
import Widget from "../../components/widget/Widget"
import "./home.scss"

const Home = () => {
    return (
        <>
            <Navbar />
            <div className="home">
                <Sidebar />
                <div className="homeContainer">
                    <Selector/>
                    <div className="widgets">
                        <Widget />
                        <Widget />
                        <Widget />
                    </div>

                </div>
            </div>
        </>

    )
}

export default Home