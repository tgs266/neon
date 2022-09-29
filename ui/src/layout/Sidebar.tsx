import { Button, List, ListItem, ListItemButton, ListItemIcon, ListItemText } from "@mui/material"
import { Box } from "@mui/system"
import React from "react"
import "./Sidebar.css"
import AppsIcon from '@mui/icons-material/Apps';
import InventoryIcon from '@mui/icons-material/Inventory';
import HomeIcon from '@mui/icons-material/Home';
import { Link } from "react-router-dom";
import SettingsIcon from '@mui/icons-material/Settings';

const boxSx = {
    backgroundColor: 'primary.main',
    '&:hover': {
        backgroundColor: 'primary.dark',
    },
}

export function Sidebar() {
    return <List disablePadding>
        <ListItem disablePadding>
            <ListItemButton component={Link} to="/">
                <ListItemIcon>
                    <HomeIcon />
                </ListItemIcon>
                <ListItemText primary="Home" />
            </ListItemButton>
        </ListItem>

        <ListItem disablePadding>
            <ListItemButton component={Link} to="/apps">
                <ListItemIcon>
                    <AppsIcon />
                </ListItemIcon>
                <ListItemText primary="Apps" />
            </ListItemButton>
        </ListItem>

        <ListItem disablePadding>
            <ListItemButton component={Link} to="/products">
                <ListItemIcon>
                    <InventoryIcon />
                </ListItemIcon>
                <ListItemText primary="Products" />
            </ListItemButton>
        </ListItem>

        <ListItem disablePadding>
            <ListItemButton component={Link} to="/settings">
                <ListItemIcon>
                    <SettingsIcon />
                </ListItemIcon>
                <ListItemText primary="Settings" />
            </ListItemButton>
        </ListItem>
    </List>
}