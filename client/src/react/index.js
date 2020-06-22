import React from 'react'

export const renderHelmetTags = (tags) => {
    return tags.map(({ id, type, attributes, value }) => {
        attributes.key = id;
        return React.createElement(type, attributes, value);
    });
}