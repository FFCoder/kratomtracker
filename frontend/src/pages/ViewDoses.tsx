import {List, ListItemButton, ListItemText } from "@mui/material";
import {getAPIServerBaseURL} from "../utils/getEnv.ts";
import {useEffect, useState} from "react";
import Typography from "@mui/material/Typography";

type Dose = {
    id: number,
    date_taken: string,
}

function parseDate(date: string): string {
    const dateObj = new Date(date);
    return dateObj.toLocaleTimeString("en-US", {timeZone: "+00:00"})
}

function ViewDoses() {
    const [loading, setLoading] = useState<boolean>(true);
    const [data, setData] = useState<Array<Dose>>([]);
    const getDoses = async () => {
        // Get doses from the API
        let id = 0;
        try {
            const controller = new AbortController();
            id = setTimeout(() => controller.abort(), 3 * 1000);
            const url = getAPIServerBaseURL() + "/doses/today";
            const response = await fetch(url, {
                method: 'GET',
                signal: controller.signal
            });
            const dataRespose = await response.json() as Array<Dose>;
            setData(dataRespose);



        } catch (err : unknown) {
            console.error(err)
        } finally {
            clearTimeout(id)
            setLoading(false);
        }
    }

    useEffect(() => {
        getDoses();
    }, []);
    return (
        <>
            <Typography variant="h5" component="h5" gutterBottom>
                Doses taken today
            </Typography>
            {loading ? "Loading..." : <List
                sx={{ width: '100%', bgcolor: 'background.paper' }}
                component="nav"
                aria-labelledby="nested-list-subheader"
            >
                {data.map(dose => (
                    <ListItemButton key={dose.id}>
                        <ListItemText primary={parseDate(dose.date_taken)} />
                    </ListItemButton>
                ))}
            </List>}
        </>
    );
}

export default ViewDoses;