
import './App.css'
import ApplicationBar from "./components/ApplicationBar.tsx";
import { useState} from "react";
import {Alert, AlertTitle, Fade } from "@mui/material";
import NavigationTabBar from "./components/NavigationTabBar.tsx";
import ViewDoses from "./pages/ViewDoses.tsx";
import ButtonPage from "./pages/ButtonPage.tsx";
import ManualAddDose from "./pages/ManualAddDose.tsx";

function App() {
    const [isSuccessful, setIsSuccessful] = useState<boolean | null>(null);

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
                Component: <ButtonPage isSuccessful={isSuccessful} setIsSuccessful={setIsSuccessful} />,
            },
            {
                Title: "Manual Add Dose",
                Component: <ManualAddDose />,
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
