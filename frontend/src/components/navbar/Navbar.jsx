import './navbar.scss'

const Navbar = () => {
    return (
        <div className="navbar">
            <div className="left">
                <span className="logo">MWA</span>
                <hr />
            </div>
            <div className="right">
                <span className="title">Maintenance Web Apps</span>
                <hr />
            </div>

        </div>
    )
}

export default Navbar