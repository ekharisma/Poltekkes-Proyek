import React from 'react';
import Chart from '../../components/chart/Chart';
import Navbar from "../../components/navbar/Navbar"
import Selector from '../../components/selector/Selector';
// import Selector from "../../components/selector/Selector"
import Sidebar from "../../components/sidebar/Sidebar"
import Widget from "../../components/widget/Widget"
// import Chart from "../../components/chart/Chart"
import "./home.css"

const Home = () => {
    return (
        <>
            <Navbar />
            <div className="home">
                <Sidebar />
                <div className="homeContainer">
                    <p className='homeLabel'>Select month</p>
                    <Selector />
                    <p className='homeLabel mt'>This month data summary</p>
                    <div className="widgets">
                        <Widget title="Maximum" suhu="15"/>
                        <Widget title="Average" suhu="7"/>
                        <Widget title="Minimum" suhu="5"/>
                    </div>
                    <p className='homeLabel mt'>Realtime data update</p>
                    <div>
                        <Chart />
                    </div>


                </div>
            </div>
        </>

    )
}

export default Home