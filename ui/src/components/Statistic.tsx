import { Card, CardContent, Typography } from "@mui/material";
import { Box, SxProps, Theme } from "@mui/system";
import React from "react";

export interface StatisticProps {
    label: string
    value: React.ReactNode
    sx?: SxProps<Theme>
}


export default function Statistic(props: StatisticProps) {
    return <Box sx={props.sx}>
        <Typography color="text.secondary" sx={{fontSize: 14}}>
            {props.label}
        </Typography>
        <Typography color="text.primary" sx={{fontSize: 16}}>
            {props.value}
        </Typography>
    </Box>
}