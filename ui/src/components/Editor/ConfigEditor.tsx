import React from "react";
import Editor from "@monaco-editor/react";

export default function ConfigEditor() {
  return (
    <Editor
      height="100%"
      defaultLanguage="yaml"
      defaultValue="# some comment"
    />
  );
}
