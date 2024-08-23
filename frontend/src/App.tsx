
import './App.css'
import ApplicationBar from "./components/ApplicationBar.tsx";
import {LogDoseButton} from "./components/LogDoseButton.tsx";
import {useEffect, useState} from "react";
import {Alert, AlertTitle, Container, Fade, Stack} from "@mui/material";
import NavigationTabBar from "./components/NavigationTabBar.tsx";
import ViewDoses from "./pages/ViewDoses.tsx";

function App() {
    const [isSuccessful, setIsSuccessful] = useState<boolean | null>(null);
    const [message, setMessage] = useState<string>('');

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

    const ButtonPage = () => {
        return (
            <>

                <Container maxWidth={"sm"} >

                    <Stack spacing={2}>
                        <LogDoseButton setIsSuccessful={setIsSuccessful} />
                        <span style={{}}>{message}</span>
                    </Stack>
                </Container>
            </>
        );
    }

  return (
    <>
        <Fade style={{display: isSuccessful ? "block" : 'none'}} in={isSuccessful !== null}>
            <Alert variant="filled" severity={isSuccessful ? "success" : "error"}>
                <AlertTitle>
                    {isSuccessful == true ? "Success" : "Error"}
                </AlertTitle>
                {isSuccessful == true ? "Dose was logged successfully!" : "Dose failed to log properly."}
            </Alert>
        </Fade>
      <ApplicationBar />

        <NavigationTabBar tabs={[
            {
                Title: "Log a Dose",
                Component: ButtonPage(),
            },
            {
                Title: "View Doses",
                Component: <ViewDoses />,
            }
        ]} />


    </>
  )
}

export default App
