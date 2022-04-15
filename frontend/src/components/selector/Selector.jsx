import React from 'react';
import TextField from '@mui/material/TextField';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';
import { AdapterMoment } from '@mui/x-date-pickers/AdapterMoment';
import './selector.scss'
const Selector = () => {
    const [value, setValue] = React.useState(null);


    return (
        // <div className='selector'>
        //     <label htmlFor="selectMonth">Select month</label>
        //     <input type="month" id='selectMonth' name='selectMonth' placeholder='select month' />
        // </div>
        <LocalizationProvider dateAdapter={AdapterMoment}>
            <DatePicker
                label="Basic example"
                value={value}
                onChange={(newValue) => {
                    setValue(newValue);
                }}
                renderInput={(params) => <TextField {...params} />}
            />
        </LocalizationProvider>
    )
}

export default Selector