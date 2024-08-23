import * as React from 'react';
import { useTheme } from '@mui/material/styles';
import SwipeableViews from 'react-swipeable-views';
import AppBar from '@mui/material/AppBar';
import Tabs from '@mui/material/Tabs';
import Tab from '@mui/material/Tab';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';

interface TabPanelProps {
    children?: React.ReactNode;
    dir?: string;
    index: number;
    value: number;
}

function TabPanel(props: TabPanelProps) {
    const { children, value, index, ...other } = props;

    return (
        <div
            role="tabpanel"
            hidden={value !== index}
            id={`full-width-tabpanel-${index}`}
            aria-labelledby={`full-width-tab-${index}`}
            {...other}
        >
            {value === index && (
                <Box sx={{ p: 3 }}>
                    <Typography>{children}</Typography>
                </Box>
            )}
        </div>
    );
}

function a11yProps(index: number) {
    return {
        id: `full-width-tab-${index}`,
        'aria-controls': `full-width-tabpanel-${index}`,
    };
}

type TabProp = {
    Title: string
    Component: JSX.Element
}

type NavigationTabBarProps = {
    tabs: Array<TabProp>
}



export default function NavigationTabBar({tabs}: NavigationTabBarProps) {
    const theme = useTheme();
    const [value, setValue] = React.useState(0);

    const handleChange = (_event: React.SyntheticEvent, newValue: number) => {
        setValue(newValue);
    };

    const handleChangeIndex = (index: number) => {
        setValue(index);
    };

    return (
        <>
            <AppBar position="static">
                <Tabs
                    value={value}
                    onChange={handleChange}
                    indicatorColor="secondary"
                    textColor="inherit"
                    variant="fullWidth"
                    aria-label="full width tabs example"
                >
                    {tabs.map((t, i) => {
                        return (
                            <Tab label={t.Title} {...a11yProps(i)} />
                        )
                    })}
                </Tabs>
            </AppBar>
            <SwipeableViews
                axis={theme.direction === 'rtl' ? 'x-reverse' : 'x'}
                index={value}
                onChangeIndex={handleChangeIndex}
            >
                {tabs.map((t, i) => {
                    return (
                        <TabPanel value={value} index={i} dir={theme.direction}>
                            {t.Component}
                        </TabPanel>
                    )
                })}
                // <TabPanel value={value} index={0} dir={theme.direction}>
                //     Item One
                // </TabPanel>
                // <TabPanel value={value} index={1} dir={theme.direction}>
                //     Item Two
                // </TabPanel>
            </SwipeableViews>
        </>
    );
}