import {Container, Stack} from "@mui/material";
import {LogDoseButton} from "../components/LogDoseButton.tsx";
import {useEffect, useState} from "react";
import {getAPIServerBaseURL} from "../utils/getEnv.ts";
import Typography from "@mui/material/Typography";

type ButtonPageProps = {
    isSuccessful: boolean | null;
    setIsSuccessful: (isSuccessful: boolean | null) => void;
}

const ButtonPage = ({isSuccessful, setIsSuccessful}: ButtonPageProps) => {
    const [message, setMessage] = useState<string>('');
    const [nextDoseTime, setNextDoseTime] = useState<Date | null>(null);
    const formatNextDoseTime = (nextDoseTime: Date | null) => {
        if (nextDoseTime) {
            return nextDoseTime.toLocaleTimeString("en-US", { timeZone: "-00:00" });
        }
        return "Loading...";
    }
    const getNextDoseTime = async () => {
        try {
            const response = await fetch(getAPIServerBaseURL() + "/doses/next")
            if (response.ok) {
                const data = await response.json();
                setNextDoseTime(new Date(data.nextDoseTime));
            }
        } catch (e) {
            console.error(e);
        }
    }
    useEffect(() => {
        if (isSuccessful === true) {
            setMessage('Dose logged successfully');
        } else if (isSuccessful === false){
            setMessage('Failed to log dose');
        } else {
            setMessage('');
        }
        setTimeout(() => setIsSuccessful(null), 5 * 1000)
    }, [isSuccessful]);

    useEffect(() => {
        getNextDoseTime();
    }, [isSuccessful]);
    return (
        <>

            <Container maxWidth={"sm"} >
                <Stack spacing={2}>
                    <LogDoseButton setIsSuccessful={setIsSuccessful} />
                    <span style={{}}>{message}</span>
                    <Typography variant={"h6"}>Next Dose Time: {formatNextDoseTime(nextDoseTime)}</Typography>
                </Stack>
            </Container>
        </>
    );
}

export default ButtonPage;