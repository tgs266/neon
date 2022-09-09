import { Collapse, List, ListItemButton, ListItemButtonProps, ListItemIcon, ListItemText } from "@mui/material";
import React, { useState } from "react";
import ExpandLess from '@mui/icons-material/ExpandLess';
import ExpandMore from '@mui/icons-material/ExpandMore';

export interface ExpandableListProps {
    primary: string
    children: React.ReactNode
}

export default function ExpandableList(props: ExpandableListProps) {
    const [open, setOpen] = useState(false)

    const { primary, children } = props

    return <>
        <ListItemButton onClick={() => setOpen(!open)}>
            <ListItemText primary={primary} />
            {open ? <ExpandLess /> : <ExpandMore />}
        </ListItemButton>
        <Collapse in={open} timeout="auto" unmountOnExit>
            <List component="div" disablePadding>
                {children}
            </List>
        </Collapse>
    </>
}