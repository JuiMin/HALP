// Define a central location for styles
// We can keep thematic elements the same using this
import { StyleSheet } from 'react-native';

// Default Thematic Coloring so you can use it in multiple objects
import Theme from './Theme';

// Generate the stylesheet
export default StyleSheet.create({
   // Define Component Specific Styling
   home: {
      flex: 1,
      backgroundColor: Theme.colors.primaryBackgroundColor,
      alignItems: 'center',
      justifyContent: 'center',
   },

   login: {
      flex: 1,
      backgroundColor: Theme.colors.primaryBackgroundColor,
      alignItems: 'center',
      justifyContent: 'center',
   },

   signup: {
      flex: 1,
      backgroundColor: Theme.colors.primaryBackgroundColor,
      alignItems: 'center',
      justifyContent: 'center',
   },

   // Navigation Bar from the default view
   navigationBar: {
      height: 49,
      flexDirection: 'row',
      borderTopWidth: StyleSheet.hairlineWidth,
      borderTopColor: 'rgba(0, 0, 0, .4)',
      backgroundColor: '#FFFFFF',
   },

   // Navigation Tabs
   navigationTab: {
      flex: 1,
      alignItems: 'center',
      justifyContent: 'center',
   },

   searchScreen: {
      backgroundColor: Theme.colors.primaryBackgroundColor
   },

   // Search bar
   searchBar: {
      backgroundColor: Theme.colors.primaryBackgroundColor,
      height: 49,
      width: "100%",
      marginBottom: 10,
      borderBottomColor: Theme.colors.primaryBackgroundColor
   },

   searchList: {
      backgroundColor: Theme.colors.primaryBackgroundColor,
      marginTop: 0
   },

   searchListItem: {
      marginTop: 0,
      backgroundColor: Theme.colors.primaryBackgroundColor,
      borderColor: Theme.colors.primaryBackgroundColor
   },

   searchTitle: {
      width: "80%",
      borderBottomColor: Theme.colors.primaryBackgroundColor
   }
});