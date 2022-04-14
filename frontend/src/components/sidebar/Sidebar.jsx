import './sidebar.scss'
import DashboardIcon from '@mui/icons-material/Dashboard';
import NotificationsActiveIcon from '@mui/icons-material/NotificationsActive';

const Sidebar = () => {
    return (
        <div className="sidebar">
            <div className="list-menu">
                <ul>
                    <li>
                        <DashboardIcon className='icon'/>
                        <br />
                        <span>Home</span>
                    </li>
                    <li>
                        <NotificationsActiveIcon className='icon'/>
                        <br />
                        <span>Notify</span>
                    </li>
                </ul>
            </div>
        </div>
    )
}

export default Sidebar