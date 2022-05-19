import React from 'react';
import './widget.css'

const Widget = (props) => {
    return (
        <div className='widget'>
            <div className="content">
                {props.title}
                <h1>{props.suhu}°C</h1>
            </div>
        </div>
    )
}

export default Widget