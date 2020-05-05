import React from "react";
import { Helmet } from "react-helmet";

const HotStoneContext = React.createContext([]);

class HotStoneWrapper extends React.Component {
    constructor(props) {
        super(props);

        const { tags } = props;
        this.state = { tags };
    }

    render() {
        const { tags } = this.state;
        const tagElements = tags.map(({ id, type, attributes, value }) => {
            attributes.key = id;
            return React.createElement(type, attributes, value);
        });
        return (
            <div>
                <Helmet>{tagElements}</Helmet>
                <HotStoneContext.Provider value={tags}>
                    {this.props.children}
                </HotStoneContext.Provider>
            </div>
        );
    }
}

const HotStone = HotStoneWrapper;

export { HotStone, HotStoneContext };