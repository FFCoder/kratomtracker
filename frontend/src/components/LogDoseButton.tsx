import {Button} from "@mui/material";
import {getAPIServerBaseURL} from "../utils/getEnv.ts";

type LogDoseButtonProps = {
    setIsSuccessful: (isSuccessful: boolean) => void;
}
export function LogDoseButton({setIsSuccessful}: LogDoseButtonProps) {
    const logDose = async () => {
        let id = 0;
        try {
            const controller = new AbortController();
            id = setTimeout(() => controller.abort(), 3 * 1000);
            const url = getAPIServerBaseURL() + "/doses/now";
            const response = await fetch(url, {
                method: 'POST',
                signal: controller.signal
            });
            setIsSuccessful(response.ok);
            if (!response.ok) {
                console.error('Failed to log dose');
            }
        } catch (err : unknown) {
            setIsSuccessful(false)
            console.error(err)
        } finally {
            clearTimeout(id)
        }
    }
    return (
        <Button variant="contained" color="secondary" onClick={logDose} size={"large"}>
            Log Dose
        </Button>
    );
}