import React from 'react';
import './App.css';
import { ApolloProvider, ApolloClient, InMemoryCache } from "@apollo/client";
import * as constants from '../../constants'

import { ChoreList } from '../ChoreList/ChoreList'

const client = new ApolloClient({
  uri: constants.HOSTNAME + "/query",
  cache: new InMemoryCache()
});

function App() {

  return (
    <div>
      <ApolloProvider client={client}>
        <ChoreList />
      </ApolloProvider>
    </div>
  );
}

export default App;
