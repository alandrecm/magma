/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @format
 */

import {
  addDecorator,
  addParameters,
} from '@storybook/react/dist/client/preview';
import {BrowserRouter} from 'react-router-dom';
import {configure} from '@storybook/react';
import {MuiThemeProvider} from '@material-ui/core/styles';
import {SnackbarProvider} from 'notistack';
import defaultTheme from '@fbcnms/ui/theme/default';
import MuiStylesThemeProvider from '@material-ui/styles/ThemeProvider';
import React from 'react';
import {themes} from '@storybook/theming';
import Theme from '../theme/symphony';

// automatically import all files ending in *.stories.js
const req = require.context('../stories', true, /.stories.js$/);
function loadStories() {
  req.keys().forEach(filename => req(filename));
}

addDecorator(story => (
  <BrowserRouter>
    <MuiThemeProvider theme={defaultTheme}>
      <MuiStylesThemeProvider theme={defaultTheme}>
        <SnackbarProvider
          maxSnack={3}
          autoHideDuration={10000}
          anchorOrigin={{
            vertical: 'bottom',
            horizontal: 'right',
          }}>
          {story()}
        </SnackbarProvider>
      </MuiStylesThemeProvider>
    </MuiThemeProvider>
  </BrowserRouter>
));

addParameters({
  options: {
    isFullscreen: false,
    showNav: true,
    showPanel: false,
    isToolshown: true,
    theme: {
      ...themes.light,
      appContentBg: Theme.palette.D10,
    },
    hierarchySeparator: /\//,
  },
});

configure(loadStories, module);
