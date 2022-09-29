import React from "react";
import Editor from "@monaco-editor/react";

export default function ConfigEditor(props: { code: string, onChange: (s: string) => void }) {
  return (
    <Editor
      height="100%"
      defaultLanguage="yaml"
      defaultValue={props.code}
      onChange={(e) => props.onChange(e)}
    />
  );
}
