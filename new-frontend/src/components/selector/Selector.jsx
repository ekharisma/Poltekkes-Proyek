import * as React from 'react';
import TextField from '@mui/material/TextField';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';

export default function Selector(props) {
    const [dateValue, setValue] = React.useState(new Date());

    return (
        <div className="selector">
            <LocalizationProvider dateAdapter={AdapterDateFns}>
                <DatePicker
                    views={['year', 'month']}
                    label="Select month"

                    value={dateValue}
                    onChange={(newValue) => {
                        setValue(newValue);

                        //console.log(Date.parse(value));
                        // console.log(value.getMonth(), value.getFullYear());
                    }}
                    renderInput={(params) => <TextField {...params} />}
                />
            </LocalizationProvider>
            {/* <h1>{dateValue.getMonth()}</h1>
            <h1>{dateValue.getFullYear()}</h1> */}
        </div>
    );
}
