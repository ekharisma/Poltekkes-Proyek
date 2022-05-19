import * as React from 'react';
import TextField from '@mui/material/TextField';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';

export default function Selector() {
    const [value, setValue] = React.useState(null);

    return (
        <div className="selector">
            <LocalizationProvider dateAdapter={AdapterDateFns}>
                <DatePicker
                    views={['year', 'month']}
                    label="Select month"
                    value={value}
                    onChange={(newValue) => {
                        setValue(newValue);
                    }}
                    renderInput={(params) => <TextField {...params} />}
                />
            </LocalizationProvider>
        </div>

    );
}
