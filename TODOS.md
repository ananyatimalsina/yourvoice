Optimizations:
Look into adding component rendering for the model table rows so that we transfer less data over the wire on Create and Edit operations.
Constraints:
- Search query parameters not aligning with model on add
- Custom filtering and components not reflecting upon add
- Dynamic add based on current page state
- Correct no items behavior on delete

Security:
Make sure that sensitive fields cannot be edited, even when tried manually

Polish:
- Use custom popup dialog for delete
- Implement error dialogs for network requests
