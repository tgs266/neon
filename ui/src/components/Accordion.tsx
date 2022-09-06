import React from "react"
import { AccordionDetails, AccordionSummary, Accordion as Ac, Typography } from '@mui/material';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import { SxProps, Theme } from "@mui/system";


export default function Accordion(props: { children?: React.ReactNode, title: React.ReactNode, defaultExpanded?: boolean, sx?: SxProps<Theme> }) {
    const { children, title, defaultExpanded } = props
    return <Ac sx={props.sx} defaultExpanded={defaultExpanded}>
        <AccordionSummary
            expandIcon={<ExpandMoreIcon />}
        >
            <Typography>{title}</Typography>
        </AccordionSummary>
        <AccordionDetails sx={{ m: 0, p: 0 }} style={{ margin: 0 }}>
            {children}
        </AccordionDetails>
    </Ac>
}