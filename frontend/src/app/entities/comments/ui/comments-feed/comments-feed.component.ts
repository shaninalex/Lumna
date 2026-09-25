import { Component } from '@angular/core';

@Component({
    selector: 'lu-comments-feed',

    template: `
        <h5 class="mb-3">Comments</h5>
        <div class="mb-4">
            <div class="d-flex justify-content-between">
                <strong>Alex</strong>
                <small class="text-muted"> Today 09:42 </small>
            </div>

            <p class="mb-0 mt-2">Initial implementation is complete. Waiting for backend review.</p>
        </div>
        <div class="mb-4">
            <div class="d-flex justify-content-between">
                <strong>Kate</strong>
                <small class="text-muted"> Yesterday </small>
            </div>
            <p class="mb-0 mt-2">Design approved.</p>
        </div>
        <textarea class="form-control" rows="4" placeholder="Write a comment..."></textarea>
        <div class="mt-3">
            <button class="btn btn-primary" disabled>Add Comment</button>
        </div>
    `,
})
export class CommentsFeedComponent {}
