import Navbar from "../../components/navbar/Navbar"
import Sidebar from "../../components/sidebar/Sidebar"
import "./home.scss"

const Home = () => {
    return (
        <>
            <Navbar />
            <div className="home">
                <Sidebar />
                <div className="homeContainer">container</div>
            </div>
        </>

    )
}

export default Home