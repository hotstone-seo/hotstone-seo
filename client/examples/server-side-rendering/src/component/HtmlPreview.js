import React, { useEffect, useState } from 'react'
import { UnControlled as CodeMirror } from 'react-codemirror2'
if (__isBrowser__) {
    require('codemirror/mode/htmlmixed/htmlmixed');
}

export default function HtmlPreview({ url }) {
    const [rawHtml, setRawHtml] = useState('')

    useEffect(() => {
        if (__isBrowser__ && url != '') {
            fetch(url).then((resp) => {
                return resp.text()
            }).then((text) => {
                //   console.log("RESP HTML: ", await resp.text())
                setRawHtml(text)
            })
        }
    }, [url]);

    return (
        <CodeMirror
            value={rawHtml}
            options={{
                mode: 'htmlmixed',
                lineNumbers: true
            }}
            onChange={(editor, data, value) => {
            }}
        />

    )
}