import {DateTimePicker, LocalizationProvider} from '@mui/x-date-pickers';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs'
import {useEffect, useState} from "react";
import {Dayjs} from "dayjs";
import {getAPIServerBaseURL} from "../utils/getEnv.ts";
import {Button, Stack} from "@mui/material";
import Typography from "@mui/material/Typography";

function ManualAddDose() {
    const [dateChosen, setDateChosen] = useState<Dayjs | null>(null);
    const [formattedDate, setFormattedDate] = useState<string | null>(null);
    useEffect(() => {
        if (dateChosen) {
            setFormattedDate(dateChosen.format("YYYY-MM-DD HH:mm:ss"));
        }
    }, [dateChosen]);
    const sendManualDose = async () => {
        if (formattedDate) {
            try {
                const url = getAPIServerBaseURL() + "/doses";
                const response = await fetch(url, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({date_taken: formattedDate})
                });
                if (response.ok) {
                    console.log("Dose logged successfully");
                } else {
                    console.error("Failed to log dose");
                }
            } catch (e) {
                console.error(e);
            }
        }
    }
    return (
       <>
           <Typography variant="h5" component="h5" gutterBottom>
                Manually Add Dose
              </Typography>
           <Typography variant={"body1"}>Choose the date and time of the dose you want to log</Typography>
           <LocalizationProvider dateAdapter={AdapterDayjs}>
                <Stack spacing={2}>
                    <DateTimePicker value={dateChosen} onChange={setDateChosen} />
                    <Button variant={"contained"} onClick={sendManualDose}>Log dose</Button>
                </Stack>
           </LocalizationProvider>
       </>
    );
}

export default ManualAddDose;